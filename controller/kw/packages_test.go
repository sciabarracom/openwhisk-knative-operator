package kw

import "github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations/packages"

func ExampleGetAllPackages() {
	SysSh("@rm -Rf /tmp/kwtest")
	Manager = NewFolderManager("kwhisk")
	params := packages.GetAllPackagesParams{}
	dump(GetAllPackages(params, nil))
	Manager.UpdatePackage("hello")
	grep(`Name`, GetAllPackages(params, nil))
	Manager.UpdatePackage("world")
	grep(`Name:`, GetAllPackages(params, nil))
	// Output:
	// (*packages.GetAllPackagesOK)({
	//  Payload: ([]*models.Package) {
	//  }
	// })
	// Name: (*string)((len=5) "hello"),
	// Namespace: (*string)((len=6) "kwhisk"),
	// Name: (*string)((len=5) "hello"),
	// Name: (*string)((len=5) "world"),
}

func ExampleUpdatePackage() {
	SysSh("@rm -Rf /tmp/kwtest")
	Manager = NewFolderManager("kwhisk")
	params := packages.UpdatePackageParams{PackageName: "bad!"}
	grep(`Error:`, UpdatePackage(params, nil))
	params = packages.UpdatePackageParams{PackageName: "good"}
	grep(`Name:`, UpdatePackage(params, nil))
	params1 := packages.GetAllPackagesParams{}
	grep(`Name:`, GetAllPackages(params1, nil))
	// Output:
	// Error: (*string)((len=83) "unable to create package 'bad!': the name of the entity contains illegal characters")
	// Name: (*string)((len=4) "good"),
	// Name: (*string)((len=4) "good"),
}

func ExampleDeletePackage() {
	SysSh("@rm -Rf /tmp/kwtest")
	Manager = NewFolderManager("kwhisk")
	Manager.UpdatePackage("hello")
	Manager.UpdateAction("hello/world")
	params := packages.DeletePackageParams{PackageName: "hello"}
	grep("Error:", DeletePackage(params, nil))
	Manager.DeleteAction("hello/world")
	dump(DeletePackage(params, nil))
	grep("Error:", DeletePackage(params, nil))
	// Output:
	// Error: (*string)((len=17) "package not empty")
	// (*packages.DeletePackageOK)({
	// })
	// Error: (*string)((len=38) "the requested resource does not exists")
}
