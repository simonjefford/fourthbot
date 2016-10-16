node('linux') {
  sh 'mkdir -p src/github.com/simonjefford/fourthbot'
  def workSpace = pwd()
  env.GOPATH = "${workSpace}"

  dir('src/github.com/simonjefford/fourthbot') {
    stage('Checkout') {
      checkout scm
    }

    def goHome = tool name: 'go', type: 'com.cloudbees.jenkins.plugins.customtools.CustomTool'

    env.PATH = "${goHome}/go/bin:${env.PATH}:${env.GOPATH}/bin"
    env.GOROOT = "${goHome}/go"

    stage('Setup') {
      sh 'go get github.com/tebeka/go2xunit'
    }
    
    stage('Test') {
      sh '2>&1 go test ./... -v | go2xunit -output tests.xml'
      junit 'tests.xml'
    }
  }
}
