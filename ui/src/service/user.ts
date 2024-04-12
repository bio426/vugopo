import base from "./base"

import * as tUser from "@/type/user"

const path = "user"

export default {
    async list() {
        const res = await base.get(path)

        return res.json<{ data: tUser.User[] }>()
    },
}
