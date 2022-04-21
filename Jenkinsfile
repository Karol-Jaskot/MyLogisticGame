pipeline {
  agent { label "linux" }
  stages {
    stage("build") {
      steps {
        sh """
          docker build -t my_logistic_game .
        """
      }
    }
    stage("run") {
      steps {
        sh """
          docker run --rm my_logistic_game
        """
      }
    }
  }
}