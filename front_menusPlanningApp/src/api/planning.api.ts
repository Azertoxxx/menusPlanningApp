import { Fetcher } from "openapi-typescript-fetch";
import type { Planning } from "../models/planning.model";

const fetcher = Fetcher.for<any>();
let isConfigured = false;

async function configureFetcher() {
    if (isConfigured) return;
    fetcher.configure({
        baseUrl: "http://localhost:8080",
        use: [
            (url, init, next) => {
                init.headers.set("Content-Type", "application/json");
                return next(url, init);
            }
        ]
    });
    isConfigured = true;
}

export async function getPlannings() {
    await configureFetcher();
    const API = fetcher.path('/plannings').method('get').create();

    return (await API({})).data as Planning[];
}

export async function createPlanning(planning: Planning) {
    await configureFetcher();
    const API = fetcher.path('/plannings').method('post').create({});

    return (await API({...planning})).data as Planning;
}

export async function updatePlanning(planning: Planning) {
    await configureFetcher();
    const API = fetcher.path(`/plannings/${planning.id}`).method('put').create({});

    return (await API({ ...planning })).data as Planning;
}

export async function deletePlanning(id: string) {
    await configureFetcher();
    const API = fetcher.path('/plannings/{id}').method('delete').create(id);

    return (await API({ id })).data as Planning;
}