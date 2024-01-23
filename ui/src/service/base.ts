import ky from "ky"

import config from "@/config"
import router from "@/router"
import useAuthStore from "@/store/auth"

const base = ky.create({
	prefixUrl: config.API_URL,
	credentials: "include",
	hooks: {
		beforeError: [
			(err) => {
				if (err.response.status == 401) {
					const authStore = useAuthStore()
					authStore.user = undefined
					setTimeout(() => alert("auth failed"), 0)
					router.push({ name: "login" })
				}
				return err
			},
		],
	},
})

export default base
