import z from "zod";

export const RoomGenerationSchema = z.object({
    roomCategoryId: z.number().positive("Room Category ID must be a positive number"),
    startDate: z.string().datetime(),
    endDate: z.string().datetime(),
    scheduledType: z.enum(["immediate", "scheduled"]),
    scheduledAt: z.string().datetime().optional(),
    priceOverride: z.number().optional()

})

export const RoomGenerationJobSchema = z.object({
    roomCategoryId: z.number().positive(),
    startDate: z.string().datetime(),
    endDate: z.string().datetime(),
    priceOverride: z.number().optional(),
    batchSize: z.number().positive().default(10)
})

