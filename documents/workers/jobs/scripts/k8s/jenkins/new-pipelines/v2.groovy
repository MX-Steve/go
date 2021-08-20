pipeline {
    agent {label 'jenkins-slave' } 
    stages {
        stage('deleteDir'){
          steps {
            sh "rm -rf ${env.WORKSPACE}/*"
          }
        }
        stage('checkout') {
            steps {
                checkout([$class: 'GitSCM', branches: [[name: params.GitProjectBranch]], doGenerateSubmoduleConfigurations: false, extensions: [[$class: 'SubmoduleOption', disableSubmodules: false, parentCredentials: true, recursiveSubmodules: true, reference: '', trackingSubmodules: false]], submoduleCfg: [], userRemoteConfigs: [[credentialsId: 'a5bff64c-58d7-4784-9ee9-336505be05da', url: "ssh://git@code.corp.alt-chain.io/source/${params.GitProjectName}.git"]]])
            }
        }
        stage('build') {
            steps {
                script {
                    withCredentials([[
                    $class: "AmazonWebServicesCredentialsBinding",
                    credentialsId: "9fd83c05-b340-4115-80ef-1468ebe5c1a7"]]) {
                        def loginAwsEcrInfo = sh(returnStdout: true, script: 'aws ecr get-login-password --region us-west-2').trim()
                        sh "docker login -u AWS -p ${loginAwsEcrInfo} https://313557088125.dkr.ecr.us-west-2.amazonaws.com"
                    }
                }
                script {
                    docker.withRegistry("https://313557088125.dkr.ecr.us-west-2.amazonaws.com/${params.NAMESPACE}/${params.ImageName}", "ecr:us-west-2:9fd83c05-b340-4115-80ef-1468ebe5c1a7") {
                        def image = docker.build("${params.NAMESPACE}/${params.ImageName}:${params.ImageTags.split(',')[0]}", "-f deploy/Dockerfile.${params.ImageTags.split(',')[0]} .")
                        for (def tag : params.ImageTags.split(',')) {
                            image.push(tag)
                        }
                    }
                }
            }
        }
        stage('deploy') {
          steps {
            sh "k8s /root/update/projects/update.sh ${params.ImageTags.split(',')[0]}"
          }
        }
        stage('cleanup') {
            steps {
                sh "docker rmi 313557088125.dkr.ecr.us-west-2.amazonaws.com/${params.NAMESPACE}/${params.ImageName}:${params.ImageTags.split(',')[0]}"
                sh "docker rmi ${params.NAMESPACE}/${params.ImageName}:${params.ImageTags.split(',')[0]}"
            }
        }
    }
}