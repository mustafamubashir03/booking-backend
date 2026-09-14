import z from "zod";

export const createRoomCategorySchema = z.object({
    hotelId: z.number().int().positive({
        message: "Hotel ID must be a positive int,eger"
    }),
    roomType: z.enum(['SINGLE',
        'DOUBLE',
        'TRIPLE',
        'DELUXE',
        'SUITE',
        'EXECUTIVE',
        'PENTHOUSE',]),
    roomCount: z.number().int().positive({
        message: "Room count must be a positive integer"
    }),
    price: z.number().positive({
        message: "Price must be a positive number"
    }),
    capacity: z.number().int().positive({
        message: "Capacity must be a positive integer"
    })
})