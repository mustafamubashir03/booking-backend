import { createRoomCategoryService } from "../services/roomCategory.service"
import { Response, Request } from "express"
import { StatusCodes } from "http-status-codes"

export async function roomCategoryController(req: Request, res: Response) {
    const roomCategory = await createRoomCategoryService(req.body)
    return res.status(StatusCodes.CREATED).json({
        message: "Room Category created successfully",
        success: true,
        data: roomCategory
    })
}