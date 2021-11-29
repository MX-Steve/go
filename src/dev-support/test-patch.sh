#!/bin/bash

ROOT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )"/.. && pwd )"
GOCACHE=/go-cache
GOPATH="/go"
export GOCACHE GOPATH

function check_go_fmt {
    result=`gofmt -l`
    if [ -n "${result}" ]; then
	    echo "Go format error within the following files:" >> $ROOT_DIR/.phabricator-comment
	    printf "%s\n" "${result}" >> $ROOT_DIR/.phabricator-comment
	    exit 1
    fi
}

function check_go_lint {
    result=`golint .`
    if [ -n "${result}" ]; then
        echo "Go lint error:" >> $ROOT_DIR/.phabricator-comment
	    printf "%s\n" "${result}" >> $ROOT_DIR/.phabricator-comment
	    exit 1
    fi
}


function check_go_vet {
    result=`go vet **/*.go 2>&1`
    if [ -n "${result}" ]; then
        echo "Go vet error:" >> $ROOT_DIR/.phabricator-comment
	    printf "%s\n" "${result}" >> $ROOT_DIR/.phabricator-comment
	    exit 1
    fi
}

function check_build {
    rm -rf $ROOT_DIR/build && mkdir -p $ROOT_DIR/build
    pushd $ROOT_DIR/build >/dev/null
    cmake -DGOPATH=${GOPATH} -GNinja .. && ninja && ctest -V
    if [ $? -ne 0 ]; then
	    echo "build failed" >> $ROOT_DIR/.phabricator-comment
	    exit 1
    fi
    popd > /dev/null
}

function check_es_lint {
    result=`git diff HEAD --name-only --diff-filter=ACMRT | grep -E "summary/.*\.(js|jsx)$"`
    if [ -n "${result}" ]; then
        pushd $ROOT_DIR/webapps/summary >/dev/null
        npm install && npm run eslint . --ext .jsx,.js
        if [ $? -ne 0 ]; then
            echo "eslint error:" >> $ROOT_DIR/.phabricator-comment
            printf "%s\n" "${result}" >> $ROOT_DIR/.phabricator-comment
            exit 1
        fi
        popd
    fi
    result=`git diff HEAD --name-only --diff-filter=ACMRT | grep -E "miningops/.*\.(js|jsx)$"`
    if [ -n "${result}" ];then
        pushd $ROOT_DIR/webapps/miningops >/dev/null
        npm install && npm run eslint . --ext .jsx,.js
        if [ $? -ne 0 ]; then
            echo "eslint error:" >> $ROOT_DIR/.phabricator-comment
            printf "%s\n" "${result}" >> $ROOT_DIR/.phabricator-comment
            exit 1
        fi
        popd
    fi
    result=`git diff HEAD --name-only --diff-filter=ACMRT | grep -E "bminer-operation/.*\.(ts|tsx)$"`
    if [ -n "${result}" ];then
        pushd $ROOT_DIR/webapps/bminer-operation >/dev/null
        npm install && npm run eslint
        if [ $? -ne 0 ]; then
            echo "eslint error:" >> $ROOT_DIR/.phabricator-comment
            printf "%s\n" "${result}" >> $ROOT_DIR/.phabricator-comment
            exit 1
        fi
        popd
    fi
    result=`git diff HEAD --name-only --diff-filter=ACMRT | grep -E "conflux-pool/.*\.(ts|tsx)$"`
    if [ -n "${result}" ];then
        pushd $ROOT_DIR/webapps/conflux-pool >/dev/null
        npm install && npm run eslint
        if [ $? -ne 0 ]; then
            echo "eslint error:" >> $ROOT_DIR/.phabricator-comment
            printf "%s\n" "${result}" >> $ROOT_DIR/.phabricator-comment
            exit 1
        fi
        popd
    fi
}

function check_npm_build {
    result=`git diff HEAD --name-only --diff-filter=ACMRT | grep -E "summary/.*\.(js|jsx)$"`
    if [ -n "${result}" ]; then
        pushd $ROOT_DIR/webapps/summary >/dev/null
        npm run build
        if [ $? -ne 0 ]; then
            echo "npm build error:" >> $ROOT_DIR/.phabricator-comment
            printf "%s\n" "${result}" >> $ROOT_DIR/.phabricator-comment
            exit 1
        fi
        popd
    fi
    result=`git diff HEAD --name-only --diff-filter=ACMRT | grep -E "miningops/.*\.(js|jsx)$"`
    if [ -n "${result}" ];then
        pushd $ROOT_DIR/webapps/miningops >/dev/null
        npm run build
        if [ $? -ne 0 ]; then
            echo "npm build error:" >> $ROOT_DIR/.phabricator-comment
            printf "%s\n" "${result}" >> $ROOT_DIR/.phabricator-comment
            exit 1
        fi
        popd
    fi
    result=`git diff HEAD --name-only --diff-filter=ACMRT | grep -E "bminer-operation/.*\.(ts|tsx)$"`
    if [ -n "${result}" ];then
        pushd $ROOT_DIR/webapps/bminer-operation >/dev/null
        npm run build
        if [ $? -ne 0 ]; then
            echo "npm build error:" >> $ROOT_DIR/.phabricator-comment
            printf "%s\n" "${result}" >> $ROOT_DIR/.phabricator-comment
            exit 1
        fi
        popd
    fi
    result=`git diff HEAD --name-only --diff-filter=ACMRT | grep -E "conflux-pool/.*\.(ts|tsx)$"`
    if [ -n "${result}" ];then
        pushd $ROOT_DIR/webapps/conflux-pool >/dev/null
        npm run build
        if [ $? -ne 0 ]; then
            echo "npm build error:" >> $ROOT_DIR/.phabricator-comment
            printf "%s\n" "${result}" >> $ROOT_DIR/.phabricator-comment
            exit 1
        fi
        popd
    fi
}

function check_webapp_clover() {
    pushd $ROOT_DIR/webapps/clover >/dev/null
    result=$(yarn install \
      && yarn run eslint --no-color \
      && yarn build)
    if [ $? -ne 0 ]; then
        echo "Clover error:" >> $ROOT_DIR/.phabricator-comment
        printf "%s\n" "${result}" >> $ROOT_DIR/.phabricator-comment
        exit 1
    fi
    popd
}

function check_python() {
    mkdir -p ${ROOT_DIR}/build/python
    python3 -m venv ${ROOT_DIR}/build/python
    source ${ROOT_DIR}/build/python/bin/activate
    pushd $ROOT_DIR/clover/fleets >/dev/null
    pip3 install -e '.[dev]'
    mkdir -p ${ROOT_DIR}/clover/fleets/backend/proto
    python setup.py gen_grpc -o ${ROOT_DIR}/clover/fleets/backend/proto
	output=$(pylint --rcfile=${ROOT_DIR}/.pylintrc --persistent=n $(git ls-files '*.py'))
    if [ $? -ne 0 ]; then
        echo "pylint error:" >> $ROOT_DIR/.phabricator-comment
        printf "%s\n" "${output}" >> $ROOT_DIR/.phabricator-comment
        exit 1
    fi
	popd
}

cat /dev/null > $ROOT_DIR/.phabricator-comment
mkdir -p $ROOT_DIR/build
cd ${ROOT_DIR}

echo "Checking Go repos"

HAS_GO=$(/usr/bin/git diff HEAD --name-only --diff-filter=ACMRT | grep -E ".*\.go$")
if [ -n "${HAS_GO}" ];then
    cd ${ROOT_DIR}/golang
    check_go_fmt
    check_go_lint
    check_build
fi

echo "Checking Python repos"
HAS_PYTHON=$(git diff HEAD --name-only --diff-filter=ACMRT | grep -E "^clover/.*\.py$")
if [ -n "${HAS_PYTHON}" ];then
    cd ${ROOT_DIR}
	check_python
fi

echo "Checking Webapps"
HAS_WEBAPP_CLOVER=$(git diff HEAD --name-only --diff-filter=ACMRT | grep -E "^webapps/clover/")
if [ -n "${HAS_WEBAPP_CLOVER}" ];then
	check_webapp_clover
fi

cd ${ROOT_DIR}
check_es_lint
check_npm_build

echo "Build succeed." > $ROOT_DIR/.phabricator-comment
