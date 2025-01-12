import { api } from "encore.dev/api";
import { hello } from "~encore/clients";

interface Response {
    greeting: string;
}

export const greeting = api(
    { expose: true, method: "GET", path: "/greeting/:name" },
    async ({ name }: { name: string }): Promise<Response> => {
        // Calling the get endpoint on the hello service
        const { message } = await hello.get();

        return { greeting: `${message} ${name}!` };
    },
);