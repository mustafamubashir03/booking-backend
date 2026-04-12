import logger from '../config/logger.config';
import Hotel from '../db/models/hotel';
import { createHotelDTO, updateHotelDTO } from '../dto/hotel.dto';
import { NotFoundError } from '../utils/errors/app.error';

export async function createHotel(hotelData: createHotelDTO) {
  try {
    const hotel = await Hotel.create({
      name: hotelData.name,
      location: hotelData.location,
      address: hotelData.address,
      rating: hotelData.rating,
      rating_count: hotelData.rating_count,
    });
    logger.info('Hotel created successfully', hotel.toJSON());
    return hotel;
  } catch {
    logger.error('Failed to create hotel');
  }
}

export async function getHotels() {
  try {
    const hotels = await Hotel.findAll({
      where: {
        deletedAt: null,
      },
    });
    return hotels;
  } catch {
    logger.error('Failed to get hotels');
  }
}
export async function getHotelById(id: number) {
  try {
    const hotel = await Hotel.findByPk(id);
    return hotel;
  } catch {
    logger.error('Failed to get hotels');
  }
}

export async function softDeleteHotelById(id: number) {
  try {
    const hotel = await Hotel.findByPk(id);
    if (!hotel) {
      logger.error('Hotel not found');
      throw new NotFoundError('No hotels found');
    }
    hotel.deletedAt = new Date();
    await hotel.save();
    logger.info('Hotel has been deleted');
    return hotel;
  } catch {
    logger.error('Failed to delete hotel');
  }
}

export async function updateHotelById(id: number, hotelData: updateHotelDTO) {
  try {
    const hotel = await Hotel.update(hotelData, {
      where: {
        id: id,
      },
    });
    return hotel;
  } catch {
    logger.error('Failed to get hotels');
  }
}
