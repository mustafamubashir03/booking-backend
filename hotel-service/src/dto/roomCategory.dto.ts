import z from "zod";
import { RoomType } from "../db/models/roomCategory";
import { RoomGenerationJobSchema } from "../validators/roomGeneration.validator";


export type createRoomCategoryDTO = {
    hotelId: number;
    roomType: RoomType;
    roomCount: number;
    price: number;
    capacity?: number;
};

export type updateRoomCategoryDTO = {
    hotelId?: number;
    roomType?: RoomType;
    roomCount?: number;
    price: number;
    capacity?: number;
};

export type RoomGenerationJobDTO = z.infer<typeof RoomGenerationJobSchema>
