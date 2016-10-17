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

    try
    {
      stage('Setup') {
        sh 'go get github.com/tebeka/go2xunit'
      }
    
      stage('Test') {
        def exit = sh(returnStatus: true, script: 'go build ./...')
        if (exit != 0) {
          error 'Build Failed'
        }
        def output = 'tests.out'
        def testResults = 'tests.xml'
        sh "2>&1 go test ./... -v | tee ${outfile}"
        exit = sh(returnStatus: true, script: "go2xunit -fail -input ${output} -output ${testResults}")
        junit testResults
      }
    }
    catch (e) {
      currentBuild.result = 'FAILED'
      throw w
    }
  }
}
