pipeline {
    agent {label 'jenkins-slave' }
    stages {
        stage('checkout') {
            steps {
                checkout([$class: 'GitSCM', branches: [[name: "platform"]], doGenerateSubmoduleConfigurations: false, extensions: [[$class: 'SubmoduleOption', disableSubmodules: false, parentCredentials: true, recursiveSubmodules: true, reference: '', trackingSubmodules: false]], submoduleCfg: [], userRemoteConfigs: [[credentialsId: 'a5bff64c-58d7-4784-9ee9-336505be05da', url: 'ssh://git@code.corp.alt-chain.io/source/mp.git']]])
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
                    try {
                        sh "sh deploy/build_node_vue.sh"
                    }
                    catch (exc) {
                        echo 'Something failed, but it is ok!'
                    }
                    docker.withRegistry("https://313557088125.dkr.ecr.us-west-2.amazonaws.com/alt-chain/reporters", "ecr:us-west-2:9fd83c05-b340-4115-80ef-1468ebe5c1a7") {
                        def image = docker.build("alt-chain/reporters:platform-frontend", "-f deploy/Dockerfile.platform .")
                    }
                    sh "docker tag alt-chain/reporters:platform-frontend https://313557088125.dkr.ecr.us-west-2.amazonaws.com/alt-chain/reporters:platform-frontend"
                    sh "docker push https://313557088125.dkr.ecr.us-west-2.amazonaws.com/alt-chain/reporters:platform-frontend"
                    sh "rm -rf *"
                }
            }
        }
    }
}