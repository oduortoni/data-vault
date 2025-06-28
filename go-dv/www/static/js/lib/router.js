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
        const segments = path.split("/").filter(Boolean);

        const route = this.#routes.find((route) => {
            const routeSegments = route.path.split("/").filter(Boolean);

            // Check if the number of segments match
            if (segments.length !== routeSegments.length) {
                return false;
            }

            return routeSegments.every((segment, index) => {
                if (segment.startsWith(":")) {
                    // This is a parameter segment, it matches any value e.g :id within posts/post/:id
                    return segments[index] !== undefined;
                }
                // This is a literal segment, must match exactly e.g post within posts/post/123
                return segment === segments[index];
            });
        });

        if (route) {
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
            route.view(app, params);
        } else {
            this.#default(app);
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