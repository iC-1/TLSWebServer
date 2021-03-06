pipeline {

  agent any
  tools { go 'Golang Plugin' }
  options { ansiColor('xterm') }

  stages {
    stage('Build') {
      steps {
        sh '''
          make build
          ls -al
        '''
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
        script {
          def branchName = BRANCH_NAME
          if (branchName == 'master') {
            branchName = 'prod'
          }
        }
        script {
          branchName = BRANCH_NAME
          if (['master'].contains(branchName)) {
            branchName = 'prod'
          }
        }
        sh """
          mkdir -p dist
          tar -zcvf dist/TLSWebServer.tar.gz TLSWebServer
          aws s3 cp dist/TLSWebServer.tar.gz s3://optimus-deploy/webserver/JenkinsBuilds/${branchName.toUpperCase()}/
        """
      }
    }
  }

  post {
    cleanup {
      cleanWs()
    }
  }
}
