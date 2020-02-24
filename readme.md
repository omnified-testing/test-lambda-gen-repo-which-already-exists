The go.mod file in the root folder indicates that this is a "Go Module".

	The go.mod file specifes the module name, which in this case is: github.com/omnified-testing/test-lambda-gen-repo-which-already-exists.
	
	The go module corresponds to the function with handle <%FUNCTIONHANDLE%>
	
	The module contains two packages:
	
	1) A utility package in utility/utility.go folder.
	2) An executable package in the main.go file
	
	The purpose of the utility package is to be able to import the function into other packages.
	
	In particular, the executable package is just a wrapper around the utility package. It just imports the utility package and executes its exported function.
	
	The executable package can be built and executed.
	