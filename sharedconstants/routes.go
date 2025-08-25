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
	NEW_ROUTE        = "/new"
	POPULAR_ROUTE    = "/popular"
	VIEWS_ROUTE      = "/views"
)

type Route struct {
	FullPath       string
	AllowedMethods []string
}

var PRIVATE_ROUTES = []Route{
	{FullPath: ME_ROUTE, AllowedMethods: []string{"GET"}},
}

var ADMIN_ROUTES = []Route{
	{FullPath: CATALOG_ROUTE + AUTHORS_ROUTE, AllowedMethods: []string{"POST"}},
	{FullPath: CATALOG_ROUTE + AUTHORS_ROUTE + "/:authorID", AllowedMethods: []string{"DELETE"}},
	{FullPath: CATALOG_ROUTE + BOOKS_ROUTE, AllowedMethods: []string{"POST"}},
	{FullPath: CATALOG_ROUTE + BOOKS_ROUTE + "/:bookID", AllowedMethods: []string{"DELETE"}},
}
