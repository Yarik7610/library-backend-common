package sharedconstants

const (
	SIGN_UP_ROUTE = "/sign-up"
	SIGN_IN_ROUTE = "/sign-in"
	ME_ROUTE      = "/me"

	CATALOG_ROUTE    = "/catalog"
	CATEGORIES_ROUTE = "/categories"
	PREVIEW_ROUTE    = "/preview"
	AUTHORS_ROUTE    = "/authors"
	BOOKS_ROUTE      = "/books"
	NEW_BOOKS_ROUTE  = "/new-books"
	SEARCH_ROUTE     = "/search"
)

type Route struct {
	Path   string
	Method string
}

var PRIVATE_ROUTES = []Route{
	{Path: ME_ROUTE, Method: "GET"},
}

var ADMIN_ROUTES = []Route{
	{Path: CATALOG_ROUTE + AUTHORS_ROUTE, Method: "POST"},
	{Path: CATALOG_ROUTE + AUTHORS_ROUTE + "/:authorID", Method: "DELETE"},
	{Path: CATALOG_ROUTE + BOOKS_ROUTE, Method: "POST"},
	{Path: CATALOG_ROUTE + BOOKS_ROUTE + "/:bookID", Method: "DELETE"},
}
