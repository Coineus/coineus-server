pipeline{

  agent any

  environment {
    dockerhub=credentials('dockerhub')
    IMAGE_NAME="safderun/coineus-server"
  }

  stages{

    stage('docker image build'){
      steps{
        sh "docker image build -t $IMAGE_NAME:latest -f ./docker/Dockerfile.server ."
      }
    }

    stage('docker image push'){
      steps{
        sh "echo $dockerhub_PSW | docker login -u $dockerhub_USR --password-stdin"
        sh "docker tag $IMAGE_NAME:latest $IMAGE_NAME:$BUILD_NUMBER"
        sh "docker push $IMAGE_NAME:latest"
        sh "docker push $IMAGE_NAME:$BUILD_NUMBER"
        sh "docker rmi $IMAGE_NAME:$BUILD_NUMBER $IMAGE_NAME:latest"
      }
    }

    stage('deploy'){
      agent {label 'ec2'}
      environment {
        SERVER_BUILD_NUMBER = "$BUILD_NUMBER"
      }
      steps{
        sh "/coineus-server/update-server.sh"
      }
    }

  
  }
}