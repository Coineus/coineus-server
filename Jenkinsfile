pipeline{

  agent any

  environment {
    IMAGE_NAME="safderun/coineus-server"
    registryCredential = 'safderun-dockerhub'
    dockerImage = ''
  }

  stages{

    stage('docker image build'){
      steps{
        echo "Building docker image"
        script{
          dockerImage = docker.build IMAGE_NAME
        }
      }
    }

    stage('docker image push'){
      steps{
        echo "Pushing docker image"
        script{
          docker.withRegistry('', registryCredential){
            dockerImage.push("$BUILD_NUMBER")
            dockerImage.push("latest")
          }
        }
        echo "Removing Docker Images"
        sh "docker rmi $IMAGE_NAME"
        sh "docker rmi $IMAGE_NAME:$BUILD_NUMBER"
        sh "docker rmi $IMAGE_NAME:latest"
      }
    }

    stage('deploy'){
      agent {label 'ec2'}
      environment {
        SERVER_BUILD_NUMBER = "$BUILD_NUMBER"
      }
      steps{
        sh "chmod +x -R ${env.WORKSPACE}"
        sh "./docker/update-server.sh"
      }
    }
  }

  post {
        success {
          mail (bcc: '', body: "Latest deploy for Coineus Server was successfull!. \n Build Number: $BUILD_NUMBER", cc: 'mlheymen.ms@gmail.com', from: 'Jenkins', replyTo: '', subject: 'Coineus Server Deploy Succesfull!', to: 'safderun@proton.me')
        }
        failure {
          mail bcc: '', body: '''Latest deploy for Coineus Server was failed!. 
          Build Number: $BUILD_NUMBER''', cc: 'mlheymen.ms@gmail.com', from: 'Jenkins', replyTo: '', subject: '!!!Coineus Server Deploy Failed!!!', to: 'safderun@proton.me'
        }
    }

}