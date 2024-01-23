const config = {
	API_URL: "http://localhost:1323/api",
	AUTH_DURATION: 120,
}

if (import.meta.env.PROD) {
	config.API_URL = "/api"
}

export default config
