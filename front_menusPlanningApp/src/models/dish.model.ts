export interface Dish {
    id?: string;
    name: string;
    imageUrl?: string;
    ingredients: string;
    description?: string;
}

export interface DishForm {
    id?: string;
    name?: string;
    imageUrl?: string;
    ingredients?: string;
    description?: string;
}

export interface GetDishesProps {
    limit?: number;
    offset?: number;
    sortBy?: string;
    sortOrder?: string;
    name?: string;
}