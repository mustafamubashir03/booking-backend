import logger from "../config/logger.config";
import Hotel from "../db/models/hotel";
import { createHotelDTO, updateHotelDTO } from "../dto/hotel.dto";

export async function createHotel(hotelData:createHotelDTO){
    try{
         const hotel = await Hotel.create({
                        name:hotelData.name,
                        location:hotelData.location,
                        address:hotelData.address,
                        rating:hotelData.rating,
                        rating_count:hotelData.rating_count
                    })
                    logger.info("Hotel created successfully",hotel.toJSON())
                    return hotel

                    }catch{
                        logger.error("Failed to create hotel")
                
                    }

}

export async function getHotels(){
    try{
        const hotels = await Hotel.findAll()
        return hotels

    }catch{
        logger.error("Failed to get hotels")

    }
}
export async function getHotelById(id:number){
    try{
        const hotel = await Hotel.findByPk(id)
        return hotel

    }catch{
        logger.error("Failed to get hotels")

    }
}

export async function deleteHotelById(id:number){
    try{
        const hotel = await Hotel.destroy({
            where:{
                id:id
            }
        })
        return hotel


    }catch{
        logger.error("Failed to delete hotel")

    }

}

export async function updateHotelById(id:number,hotelData:updateHotelDTO){
    try{
        const hotel = await Hotel.update(hotelData,{
            where:{
                id:id
            }
        })
        return hotel

    }catch{
        logger.error("Failed to get hotels")

    }

}
