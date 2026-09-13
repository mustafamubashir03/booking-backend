import { StatusCodes } from "http-status-codes"
import { generateRooms } from "../services/roomGeneration.service"
import { RoomGenerationJobDTO } from "../dto/roomCategory.dto"
import { Request, Response } from "express"
import { addRoomGenerationJobToQueue } from "../producers/roomGeneration.producer"


export async function roomGenerationController(req: Request, res: Response) {
    await addRoomGenerationJobToQueue(req?.body as RoomGenerationJobDTO)
    res.status(StatusCodes.OK).json({
        message: "Job for room generation has been added to the queue",
        data: null,
        success: true,
    })
}
