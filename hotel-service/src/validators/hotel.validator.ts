import { z } from 'zod';

export const hotelCreateSchema = z.object({
  name: z.string().min(3),
  address: z.string().min(3),
  location: z.string().min(3),
  rating: z.number().optional(),
  rating_count: z.number().optional(),
});
export const hotelUpdateSchema = z.object({
  name: z.string().min(3).optional(),
  address: z.string().min(3).optional(),
  location: z.string().min(3).optional(),
  rating: z.number().optional().optional(),
  rating_count: z.number().optional(),
});
