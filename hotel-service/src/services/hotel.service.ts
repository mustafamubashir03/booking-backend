import { createHotelDTO, updateHotelDTO } from "../dto/hotel.dto";
import { createHotel, deleteHotelById, getHotelById, getHotels, updateHotelById } from "../repositories/hotel.repository";

export async function createHotelService(hotelData:createHotelDTO){
    const hotel = await createHotel(hotelData)
    return hotel
}

export async function getHotelService(){
    const hotels = await getHotels()
    return hotels
}
export async function getHotelByIdService(id:number){
    const hotel = await getHotelById(id)
    return hotel
}

export async function updateHotelByIdService({id,hotelData}:{id:number,hotelData:updateHotelDTO}){
    const hotel = await updateHotelById(id,hotelData)
    return hotel
}

export async function deleteHotelByIdService(id:number){
    const hotel = await deleteHotelById(id)
    return hotel
}