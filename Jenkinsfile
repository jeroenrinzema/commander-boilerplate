pipeline {
  agent any
  stages {
    stage('Build images') {
      steps {
        dir(path: 'command/') {
          sh 'docker build --build-arg GO_PROJECT_PATH=github.com/sysco-middleware/commander-boilerplate/logic -t commander_command .'
        }

      }
    }
  }
}