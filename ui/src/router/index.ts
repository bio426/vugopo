import { createRouter, createWebHistory, type RouteLocationNormalized } from "vue-router"

import useAuthStore from "@/store/auth"

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: "/",
            name: "dashboard",
            component: () => import("../views/dashboard.vue"),
        },
        {
            path: "/login",
            name: "login",
            component: () => import("../views/login.vue"),
        },
        {
            path: "/task",
            name: "task",
            component: () => import("../views/task.vue"),
        },
        {
            path: "/task/create",
            name: "task-create",
            component: () => import("../views/create-task.vue"),
        },
        {
            path: "/user",
            name: "user",
            component: () => import("../views/user/list.vue"),
        },
        {
            path: "/user/create",
            name: "user-create",
            component: () => import("../views/user/create.vue"),
        },
        // Test permissions
        {
            path: "/owner",
            name: "owner",
            component: () => import("../views/owners.vue"),
            meta: { permittedRoles: ["owner"] }
        },
        {
            path: "/employee",
            name: "employee",
            component: () => import("../views/employees.vue"),
        },
        {
            path: "/anyone",
            name: "anyone",
            component: () => import("../views/anyone.vue"),
        },
        // ~
        {
            path: "/:pathMatch(.*)*",
            name: "not-found",
            component: () => import("../views/not-found.vue"),
        },
    ],
})

router.beforeEach((to, _) => {
    const authStore = useAuthStore()
    const routeName = to.name as string
    const publicRoutes = ["login", "not-found"]

    if (publicRoutes.includes(routeName)) {
        if (authStore.authenticated && routeName == "login") {
            return { name: "dashboard" }
        }
        return true
    } else {
        if (!authStore.authenticated) {
            unauthorizedAlert()
            return { name: "login" }
        }
        const userRole = authStore.user?.role as string
        const hasRolePermission = checkRolePermission(to, userRole)
        if (!hasRolePermission) {
            forbiddenAlert()
            return { name: "dashboard" }
        }
        return true
    }
})

function checkRolePermission(to: RouteLocationNormalized, role: string): boolean {
    const allowedRoles = to.meta.permittedRoles as (string[] | undefined)
    if (allowedRoles == undefined || allowedRoles.includes(role)) return true
    return false
}

function unauthorizedAlert() {
    setTimeout(() => alert("Unauthorized"), 0)
}

function forbiddenAlert() {
    setTimeout(() => alert("Forbidden"), 0)
}

export default router
