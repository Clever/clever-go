// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/Clever/clever-go/v4/restapi/operations"
	"github.com/Clever/clever-go/v4/restapi/operations/courses"
	"github.com/Clever/clever-go/v4/restapi/operations/districts"
	"github.com/Clever/clever-go/v4/restapi/operations/events"
	"github.com/Clever/clever-go/v4/restapi/operations/resources"
	"github.com/Clever/clever-go/v4/restapi/operations/schools"
	"github.com/Clever/clever-go/v4/restapi/operations/sections"
	"github.com/Clever/clever-go/v4/restapi/operations/terms"
	"github.com/Clever/clever-go/v4/restapi/operations/users"
)

//go:generate swagger generate server --target ../../clever-go --name Clever --spec ../swagger.yml --principal interface{}

func configureFlags(api *operations.CleverAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CleverAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.OauthAuth == nil {
		api.OauthAuth = func(token string, scopes []string) (interface{}, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (oauth) has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.UsersGetContactsForUserHandler == nil {
		api.UsersGetContactsForUserHandler = users.GetContactsForUserHandlerFunc(func(params users.GetContactsForUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation users.GetContactsForUser has not yet been implemented")
		})
	}
	if api.CoursesGetCourseHandler == nil {
		api.CoursesGetCourseHandler = courses.GetCourseHandlerFunc(func(params courses.GetCourseParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation courses.GetCourse has not yet been implemented")
		})
	}
	if api.SectionsGetCourseForSectionHandler == nil {
		api.SectionsGetCourseForSectionHandler = sections.GetCourseForSectionHandlerFunc(func(params sections.GetCourseForSectionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation sections.GetCourseForSection has not yet been implemented")
		})
	}
	if api.CoursesGetCoursesHandler == nil {
		api.CoursesGetCoursesHandler = courses.GetCoursesHandlerFunc(func(params courses.GetCoursesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation courses.GetCourses has not yet been implemented")
		})
	}
	if api.ResourcesGetCoursesForResourceHandler == nil {
		api.ResourcesGetCoursesForResourceHandler = resources.GetCoursesForResourceHandlerFunc(func(params resources.GetCoursesForResourceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation resources.GetCoursesForResource has not yet been implemented")
		})
	}
	if api.SchoolsGetCoursesForSchoolHandler == nil {
		api.SchoolsGetCoursesForSchoolHandler = schools.GetCoursesForSchoolHandlerFunc(func(params schools.GetCoursesForSchoolParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation schools.GetCoursesForSchool has not yet been implemented")
		})
	}
	if api.DistrictsGetDistrictHandler == nil {
		api.DistrictsGetDistrictHandler = districts.GetDistrictHandlerFunc(func(params districts.GetDistrictParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation districts.GetDistrict has not yet been implemented")
		})
	}
	if api.CoursesGetDistrictForCourseHandler == nil {
		api.CoursesGetDistrictForCourseHandler = courses.GetDistrictForCourseHandlerFunc(func(params courses.GetDistrictForCourseParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation courses.GetDistrictForCourse has not yet been implemented")
		})
	}
	if api.SchoolsGetDistrictForSchoolHandler == nil {
		api.SchoolsGetDistrictForSchoolHandler = schools.GetDistrictForSchoolHandlerFunc(func(params schools.GetDistrictForSchoolParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation schools.GetDistrictForSchool has not yet been implemented")
		})
	}
	if api.SectionsGetDistrictForSectionHandler == nil {
		api.SectionsGetDistrictForSectionHandler = sections.GetDistrictForSectionHandlerFunc(func(params sections.GetDistrictForSectionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation sections.GetDistrictForSection has not yet been implemented")
		})
	}
	if api.TermsGetDistrictForTermHandler == nil {
		api.TermsGetDistrictForTermHandler = terms.GetDistrictForTermHandlerFunc(func(params terms.GetDistrictForTermParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation terms.GetDistrictForTerm has not yet been implemented")
		})
	}
	if api.UsersGetDistrictForUserHandler == nil {
		api.UsersGetDistrictForUserHandler = users.GetDistrictForUserHandlerFunc(func(params users.GetDistrictForUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation users.GetDistrictForUser has not yet been implemented")
		})
	}
	if api.DistrictsGetDistrictsHandler == nil {
		api.DistrictsGetDistrictsHandler = districts.GetDistrictsHandlerFunc(func(params districts.GetDistrictsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation districts.GetDistricts has not yet been implemented")
		})
	}
	if api.EventsGetEventHandler == nil {
		api.EventsGetEventHandler = events.GetEventHandlerFunc(func(params events.GetEventParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation events.GetEvent has not yet been implemented")
		})
	}
	if api.EventsGetEventsHandler == nil {
		api.EventsGetEventsHandler = events.GetEventsHandlerFunc(func(params events.GetEventsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation events.GetEvents has not yet been implemented")
		})
	}
	if api.ResourcesGetResourceHandler == nil {
		api.ResourcesGetResourceHandler = resources.GetResourceHandlerFunc(func(params resources.GetResourceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation resources.GetResource has not yet been implemented")
		})
	}
	if api.ResourcesGetResourcesHandler == nil {
		api.ResourcesGetResourcesHandler = resources.GetResourcesHandlerFunc(func(params resources.GetResourcesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation resources.GetResources has not yet been implemented")
		})
	}
	if api.CoursesGetResourcesForCourseHandler == nil {
		api.CoursesGetResourcesForCourseHandler = courses.GetResourcesForCourseHandlerFunc(func(params courses.GetResourcesForCourseParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation courses.GetResourcesForCourse has not yet been implemented")
		})
	}
	if api.SectionsGetResourcesForSectionHandler == nil {
		api.SectionsGetResourcesForSectionHandler = sections.GetResourcesForSectionHandlerFunc(func(params sections.GetResourcesForSectionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation sections.GetResourcesForSection has not yet been implemented")
		})
	}
	if api.UsersGetResourcesForUserHandler == nil {
		api.UsersGetResourcesForUserHandler = users.GetResourcesForUserHandlerFunc(func(params users.GetResourcesForUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation users.GetResourcesForUser has not yet been implemented")
		})
	}
	if api.SchoolsGetSchoolHandler == nil {
		api.SchoolsGetSchoolHandler = schools.GetSchoolHandlerFunc(func(params schools.GetSchoolParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation schools.GetSchool has not yet been implemented")
		})
	}
	if api.SectionsGetSchoolForSectionHandler == nil {
		api.SectionsGetSchoolForSectionHandler = sections.GetSchoolForSectionHandlerFunc(func(params sections.GetSchoolForSectionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation sections.GetSchoolForSection has not yet been implemented")
		})
	}
	if api.SchoolsGetSchoolsHandler == nil {
		api.SchoolsGetSchoolsHandler = schools.GetSchoolsHandlerFunc(func(params schools.GetSchoolsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation schools.GetSchools has not yet been implemented")
		})
	}
	if api.CoursesGetSchoolsForCourseHandler == nil {
		api.CoursesGetSchoolsForCourseHandler = courses.GetSchoolsForCourseHandlerFunc(func(params courses.GetSchoolsForCourseParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation courses.GetSchoolsForCourse has not yet been implemented")
		})
	}
	if api.TermsGetSchoolsForTermHandler == nil {
		api.TermsGetSchoolsForTermHandler = terms.GetSchoolsForTermHandlerFunc(func(params terms.GetSchoolsForTermParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation terms.GetSchoolsForTerm has not yet been implemented")
		})
	}
	if api.UsersGetSchoolsForUserHandler == nil {
		api.UsersGetSchoolsForUserHandler = users.GetSchoolsForUserHandlerFunc(func(params users.GetSchoolsForUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation users.GetSchoolsForUser has not yet been implemented")
		})
	}
	if api.SectionsGetSectionHandler == nil {
		api.SectionsGetSectionHandler = sections.GetSectionHandlerFunc(func(params sections.GetSectionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation sections.GetSection has not yet been implemented")
		})
	}
	if api.SectionsGetSectionsHandler == nil {
		api.SectionsGetSectionsHandler = sections.GetSectionsHandlerFunc(func(params sections.GetSectionsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation sections.GetSections has not yet been implemented")
		})
	}
	if api.CoursesGetSectionsForCourseHandler == nil {
		api.CoursesGetSectionsForCourseHandler = courses.GetSectionsForCourseHandlerFunc(func(params courses.GetSectionsForCourseParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation courses.GetSectionsForCourse has not yet been implemented")
		})
	}
	if api.ResourcesGetSectionsForResourceHandler == nil {
		api.ResourcesGetSectionsForResourceHandler = resources.GetSectionsForResourceHandlerFunc(func(params resources.GetSectionsForResourceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation resources.GetSectionsForResource has not yet been implemented")
		})
	}
	if api.SchoolsGetSectionsForSchoolHandler == nil {
		api.SchoolsGetSectionsForSchoolHandler = schools.GetSectionsForSchoolHandlerFunc(func(params schools.GetSectionsForSchoolParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation schools.GetSectionsForSchool has not yet been implemented")
		})
	}
	if api.TermsGetSectionsForTermHandler == nil {
		api.TermsGetSectionsForTermHandler = terms.GetSectionsForTermHandlerFunc(func(params terms.GetSectionsForTermParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation terms.GetSectionsForTerm has not yet been implemented")
		})
	}
	if api.UsersGetSectionsForUserHandler == nil {
		api.UsersGetSectionsForUserHandler = users.GetSectionsForUserHandlerFunc(func(params users.GetSectionsForUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation users.GetSectionsForUser has not yet been implemented")
		})
	}
	if api.UsersGetStudentsForUserHandler == nil {
		api.UsersGetStudentsForUserHandler = users.GetStudentsForUserHandlerFunc(func(params users.GetStudentsForUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation users.GetStudentsForUser has not yet been implemented")
		})
	}
	if api.UsersGetTeachersForUserHandler == nil {
		api.UsersGetTeachersForUserHandler = users.GetTeachersForUserHandlerFunc(func(params users.GetTeachersForUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation users.GetTeachersForUser has not yet been implemented")
		})
	}
	if api.TermsGetTermHandler == nil {
		api.TermsGetTermHandler = terms.GetTermHandlerFunc(func(params terms.GetTermParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation terms.GetTerm has not yet been implemented")
		})
	}
	if api.SectionsGetTermForSectionHandler == nil {
		api.SectionsGetTermForSectionHandler = sections.GetTermForSectionHandlerFunc(func(params sections.GetTermForSectionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation sections.GetTermForSection has not yet been implemented")
		})
	}
	if api.TermsGetTermsHandler == nil {
		api.TermsGetTermsHandler = terms.GetTermsHandlerFunc(func(params terms.GetTermsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation terms.GetTerms has not yet been implemented")
		})
	}
	if api.SchoolsGetTermsForSchoolHandler == nil {
		api.SchoolsGetTermsForSchoolHandler = schools.GetTermsForSchoolHandlerFunc(func(params schools.GetTermsForSchoolParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation schools.GetTermsForSchool has not yet been implemented")
		})
	}
	if api.UsersGetUserHandler == nil {
		api.UsersGetUserHandler = users.GetUserHandlerFunc(func(params users.GetUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation users.GetUser has not yet been implemented")
		})
	}
	if api.UsersGetUsersHandler == nil {
		api.UsersGetUsersHandler = users.GetUsersHandlerFunc(func(params users.GetUsersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation users.GetUsers has not yet been implemented")
		})
	}
	if api.ResourcesGetUsersForResourceHandler == nil {
		api.ResourcesGetUsersForResourceHandler = resources.GetUsersForResourceHandlerFunc(func(params resources.GetUsersForResourceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation resources.GetUsersForResource has not yet been implemented")
		})
	}
	if api.SchoolsGetUsersForSchoolHandler == nil {
		api.SchoolsGetUsersForSchoolHandler = schools.GetUsersForSchoolHandlerFunc(func(params schools.GetUsersForSchoolParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation schools.GetUsersForSchool has not yet been implemented")
		})
	}
	if api.SectionsGetUsersForSectionHandler == nil {
		api.SectionsGetUsersForSectionHandler = sections.GetUsersForSectionHandlerFunc(func(params sections.GetUsersForSectionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation sections.GetUsersForSection has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
