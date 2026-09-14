import z from "zod";
import { RoomType } from "../db/models/roomCategory";
import { roomGenerationJobSchema } from "../validators/roomGeneration.validator";



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

export type roomGenerationJobDTO = z.infer<typeof roomGenerationJobSchema>
