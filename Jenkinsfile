pipeline {
    agent { node 'JenkinsAutoscaledBuildNode' }

    options { ansiColor('xterm') }

    tools { go }

    stages {
        stage('Check GO Environment') {
            steps {
                sh 'echo $GOROOT'
                sh 'echo $GOPATH'
            }
        }
        stage('Build') {
            steps {
                sh 'make build'
            }
        }
        stage('Deploy') {
            when {
                beforeAgent true;
                expression {
                    return ['dev', 'test', 'master'].contains(BRANCH_NAME);
                }
            }
            steps {
                sh 'mkdir -p dist'
                sh 'tar -zcvf dist/TLSWebServer.tar.gz TLSWebServer'
                script {
                    def branchName = BRANCH_NAME
                    if (branchName == 'master') {
                        branchName = 'prod'
                    }
                    sh "aws s3 cp dist/TLSWebServer.tar.gz s3://optimus-deploy/webserver/JenkinsBuilds/${branchName.toUpperCase()}/"
                }
            }
        }
    }
}
