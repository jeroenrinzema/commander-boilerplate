node {
  agent any
  stages {
    agent any
    def app

    stage('Clone repository') {
      checkout master
    }

    stage('Build image') {
      steps {
        dir(path: 'command/') {
          app = docker.build("commander-command")
        }
      }
    }
  }
}
