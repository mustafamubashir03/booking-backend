import { RoomType } from "../db/models/roomCategory";


export type createRoomCategoryDTO = {
    hotelId: number;
    roomType: RoomType;
    roomCount: number;
    name?: string;
    price?: number;
    capacity?: number;
};

export type updateRoomCategoryDTO = {
    hotelId?: number;
    roomType?: RoomType;
    roomCount?: number;
    name?: string;
    price?: number;
    capacity?: number;
};
