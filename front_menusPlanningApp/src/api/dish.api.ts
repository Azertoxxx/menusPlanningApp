import { Fetcher } from "openapi-typescript-fetch";
import type { Dish, GetDishesProps } from "../models/dish.model";

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

export async function getDishes(args: GetDishesProps) {
    await configureFetcher();
    const API = fetcher.path('/dishes').method('get').create();

    return (await API(args)).data as Dish[];
}

export async function getDish(id: string) {
    await configureFetcher();
    const API = fetcher.path('/dishes/{id}').method('get').create();

    return (await API({ id })).data as Dish;
}

export async function createDish(dish: Dish) {
    await configureFetcher();
    const API = fetcher.path('/dishes').method('post').create({});

    return (await API({...dish})).data as Dish;
}

export async function updateDish(dish: Dish) {
    await configureFetcher();
    const API = fetcher.path(`/dishes/${dish.id}`).method('put').create({});

    return (await API({ ...dish })).data as Dish;
}

export async function deleteDish(id: string) {
    await configureFetcher();
    const API = fetcher.path('/dishes/{id}').method('delete').create(id);

    return (await API({ id })).data as Dish;
}