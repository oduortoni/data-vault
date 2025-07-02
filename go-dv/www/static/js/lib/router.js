/*
 * file: blogsy/resources/js/lib/router.js
 * description: This file is used to route the application to the correct view based on the current path. Uses the history API to navigate between pages.
 * author: toni
 * date: 2025-06-02
 * version: 1.0.0
 * license: MIT
 * copyright: 2025 toni
 * contact: oduortoni@gmail.com
 */

class Router {
    #routes = [];
    #default = null;

    constructor() {}

    register(path, view) {
        this.#routes.push({ path, view });
    }

    route(path) {
        console.log(`Routing to: ${path}`);
        const segments = path.split("/").filter(Boolean);

        const route = this.#routes.find((route) => {
            const routeSegments = route.path.split("/").filter(Boolean);

            // Check if the number of segments match
            if (segments.length !== routeSegments.length) {
                return false;
            }

            // Use a standard for-loop for clearer, more explicit matching logic.
            for (let i = 0; i < routeSegments.length; i++) {
                const routeSegment = routeSegments[i];
                const pathSegment = segments[i];

                // If the route segment is a parameter (e.g., ":id"), it's a match.
                if (routeSegment.startsWith(":")) {
                    continue;
                }

                // If the literal segments do not match, this is not the correct route.
                if (routeSegment !== pathSegment) {
                    return false;
                }
            }

            return true; // All segments matched.
        });

        if (route) {
            console.log(`Matched route: ${route.path}`);
            // Extract parameters if any
            const params = {};
            const routeSegments = route.path.split("/").filter(Boolean);

            routeSegments.forEach((segment, index) => {
                if (segment.startsWith(":")) {
                    const paramName = segment.slice(1); // Remove the ":"
                    params[paramName] = segments[index];
                }
            });

            // Pass both app and params to the view
            route.view(params);
        } else {
            console.log("No route matched");
            this.#default();
        }
    }

    fallback(view) {
        this.#default = view;
    }

    /**
     * Pushes a new route and calls the router
     */
    navigate(path) {
        history.pushState({}, "", path);
        this.route(path);
    }

    get routes() {
        return this.#routes;
    }

    get default() {
        return this.#default;
    }
}

export default Router;