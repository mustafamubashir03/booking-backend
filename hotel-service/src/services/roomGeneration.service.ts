import { CreationAttributes } from "sequelize";
import Room from "../db/models/room";
import RoomCategory from "../db/models/roomCategory";
import { RoomGenerationJobDTO } from "../dto/roomCategory.dto";
import { RoomRepository } from "../repositories/room.repository";
import { BadRequestError, NotFoundError } from "../utils/errors/app.error";
import { roomCategoryRepository } from "./roomCategory.service";
import logger from "../config/logger.config";


export const roomRepository = new RoomRepository()

export async function generateRooms(jobData: RoomGenerationJobDTO) {
    let totalRoomsCreated = 0
    let totalDatesCovered = 0
    const roomCategory = await roomCategoryRepository.findById(jobData.roomCategoryId)
    if (!roomCategory) {
        throw new NotFoundError("Room category not found")
    }
    const startDate = new Date(jobData.startDate)
    const endDate = new Date(jobData.endDate)
    if (startDate.getTime() > endDate.getTime()) {
        throw new BadRequestError("Start date must be before end date")
    }
    if (startDate.getTime() < new Date().getTime()) {
        throw new BadRequestError("Start date must be in future")
    }
    const totalDays = Math.ceil((endDate.getTime() - startDate.getTime()) / (1000 * 60 * 60 * 24))
    logger.info(`Creating rooms for ${totalDays} days`)
    if (totalDays > jobData.batchSize) {
        throw new BadRequestError(`Total days should not be greater than batch size of ${jobData.batchSize}`)
    }
    const batchSize = jobData.batchSize || 100;
    let currentBatchStart = new Date(startDate);
    while (currentBatchStart < endDate) {
        const batchEndDate = new Date(currentBatchStart);
        batchEndDate.setDate(batchEndDate.getDate() + batchSize);
        if (batchEndDate > endDate) {
            batchEndDate.setTime(endDate.getTime())
        }
        const batchResult = await processDateBatch(roomCategory, currentBatchStart, batchEndDate, jobData.priceOverride)
        totalRoomsCreated += batchResult.roomsCreated
        totalDatesCovered += batchResult.datesProcessed
        currentBatchStart.setTime(batchEndDate.getTime())

    }
    return { totalRoomsCreated, totalDatesCovered }
}

export async function processDateBatch(roomCategory: RoomCategory, startDate: Date, endDate: Date, priceOverride: number | undefined) {
    const currentDate = new Date(startDate)
    const roomsToCreate: CreationAttributes<Room>[] = []
    let datesProcessed = 0
    let roomsCreated = 0
    const existingRooms = await roomRepository.findByRoomCategoryIdAndDateRange(roomCategory.id, startDate, endDate)
    const existingDates = new Set(
        existingRooms.map(room =>
            room.dateOfAvailability.toISOString().split("T")[0]
        )
    )
    while (currentDate <= endDate) {
        const dateKey = currentDate.toISOString().split("T")[0];

        if (!existingDates.has(dateKey)) {
            roomsToCreate.push({
                hotelId: roomCategory.hotelId,
                roomCategoryId: roomCategory.id,
                dateOfAvailability: new Date(currentDate),
                price: priceOverride !== undefined ? priceOverride : roomCategory.price,
                createdAt: new Date(),
                updatedAt: new Date(),
                deletedAt: null
            });
        }
        currentDate.setDate(currentDate.getDate() + 1);
        datesProcessed++;
    }
    if (roomsToCreate.length > 0) {
        await roomRepository.bulkCreate(roomsToCreate)
        roomsCreated += roomsToCreate.length
    }
    return { roomsCreated, datesProcessed }

}