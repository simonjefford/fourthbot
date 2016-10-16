node('linux') {
  sh 'mkdir -p src/github.com/simonjefford/fourthbot'
  def workSpace = pwd()
  env.GOPATH = "${workSpace}"

  dir('src/github.com/simonjefford/fourthbot') {
    stage('Checkout') {
      checkout scm
    }

    def goHome = tool name: 'go', type: 'com.cloudbees.jenkins.plugins.customtools.CustomTool'

    env.PATH = "${goHome}/go/bin:${env.PATH}"
    env.GOROOT = "${goHome}/go"
    
    stage('Test') {
      sh 'go test -v ./...'
    }
  }
}
