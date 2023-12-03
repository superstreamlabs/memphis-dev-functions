def versionTag
pipeline {
  environment {
    gitBranch = "${env.BRANCH_NAME}"
    gitURL = "git@github.com:Memphisdev/memphis-functions.js.git"
    repoUrlPrefix = "memphisos"
  }

  agent {
    label 'small-ec2-fleet'
  }

  stages {
    stage ('Connect GIT repository') {
      steps {
        git credentialsId: 'main-github', url: "git@github.com:Memphisdev/memphis-functions.js.git", branch: "${env.gitBranch}" 
      }
    }

    stage('Define version - BETA') {
      when {branch 'master'}
      steps {
        script {
          versionTag = readFile('./version-beta.conf')
          sh """
            sed -i -r "s/\\"memphis-functions/\\"memphis-functions-beta/g" ./package.json
          """
        }
      }
    }
    stage('Define version - LATEST') {
      when {branch 'latest'}
      steps {
        script {
          versionTag = readFile('./version.conf')
        }
      }
    }

    stage('Install NPM') {
      steps {
        sh """
          sudo dnf install https://rpm.nodesource.com/pub_18.x/nodistro/repo/nodesource-release-nodistro-1.noarch.rpm -y
          sudo dnf install nodejs -y --setopt=nodesource-nodejs.module_hotfixes=1
          sudo dnf install -y /usr/bin/g++
        """
      }
    }

    stage('Push to NPM') {
      steps {
        sh """
        sed -i -r "s/version\\": \\"[0-9].[0-9].[0-9]/version\\": \\"$versionTag/g" ./package.json
        sudo npm install
      """
        withCredentials([string(credentialsId: 'npm_token', variable: 'npm_token')]) {
          sh """
            echo //registry.npmjs.org/:_authToken=$npm_token > .npmrc
            npm publish
          """
        }
      }
    }

    stage('Checkout to version branch') {
      when {branch 'latest'}
      steps {
        sh """
          sudo dnf config-manager --add-repo https://cli.github.com/packages/rpm/gh-cli.repo -y
          sudo dnf install gh -y
          sudo dnf install jq -y
        """
        withCredentials([sshUserPrivateKey(keyFileVariable:'check',credentialsId: 'main-github')]) {
          sh """
            GIT_SSH_COMMAND='ssh -i $check'  git checkout -b $versionTag
            GIT_SSH_COMMAND='ssh -i $check' git push --set-upstream origin $versionTag
          """
        }
        withCredentials([string(credentialsId: 'gh_token', variable: 'GH_TOKEN')]) {
          sh """
            gh release create $versionTag --generate-notes
          """
        }
      }
    }

  }
  post {
    always {
      cleanWs()
    }
    success {
      notifySuccessful()
    }
    failure {
      notifyFailed()
    }
  }
}
def notifySuccessful() {
    emailext (
        subject: "SUCCESSFUL: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]'",
        body: """SUCCESSFUL: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]':
        Check console output and connection attributes at ${env.BUILD_URL}""",
        recipientProviders: [requestor()]
    )
}
def notifyFailed() {
    emailext (
        subject: "FAILED: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]'",
        body: """FAILED: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]':
        Check console output at ${env.BUILD_URL}""",
        recipientProviders: [requestor()]
    )
}