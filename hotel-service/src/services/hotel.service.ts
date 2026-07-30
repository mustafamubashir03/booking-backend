import { createHotelDTO, updateHotelDTO } from '../dto/hotel.dto';
import {
  HotelRepository
} from '../repositories/hotel.repository';


export const hotelRepository = new HotelRepository()

export async function createHotelService(hotelData: createHotelDTO) {
  const hotel = await   hotelRepository.create(hotelData);
  return hotel;
}

export async function getHotelService() {
  const hotels = await hotelRepository.findAll();
  return hotels;
}
export async function getHotelByIdService(id: number) {
  const hotel = await hotelRepository.findById(id);
  return hotel;
}

export async function updateHotelByIdService({
  id,
  hotelData,
}: {
  id: number;
  hotelData: updateHotelDTO;
}) {
  const hotel = await hotelRepository.update(id, hotelData);
  return hotel;
}

export async function deleteHotelByIdService(id: number) {
  const hotel = await hotelRepository.softDelete(id);
  return hotel;
}
