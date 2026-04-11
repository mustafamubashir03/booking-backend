
import { Request, Response } from "express";
import { createHotelService, deleteHotelByIdService, getHotelByIdService, getHotelService, updateHotelByIdService } from "../services/hotel.service";

export async function createHotelController(req:Request,res:Response){
    
    const hotelCreated = await createHotelService(req?.body)
    res.status(201).json({
            message:"Hotel created successfully",
            data:hotelCreated,
            success:true
        })
}
export async function getHotelController(req:Request,res:Response){
    
    const hotelsFound = await getHotelService()
    res.status(201).json({
            message:"Hotels found successfully",
            data:hotelsFound,
            success:true
        })
}
export async function getHotelByIdController(req:Request,res:Response){
    
    const hotelFound = await getHotelByIdService(Number(req.params.id))
    res.status(201).json({
            message:"Hotel found successfully",
            data:hotelFound,
            success:true
        })
}

export async function updateHotelByIdController(req:Request, res:Response){
    const hotelUpdated = await updateHotelByIdService({id:Number(req.params.id),hotelData:req.body})
    res.status(201).json({
        message:"Hotel updated successfully",
        data:hotelUpdated,
        success:true
    })

}

export async function deleteHotelByIdController(req:Request,res:Response){
    const hotelDeleted =  await deleteHotelByIdService(Number(req.params.id))
    res.status(201).json({
        message:"Hotel deleted successfully",
        data:hotelDeleted,
        success:true
    })
}
