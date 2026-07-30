import logger from '../config/logger.config';
import Hotel from '../db/models/hotel';
import { NotFoundError } from '../utils/errors/app.error';
import { BaseRepository } from './base.repository';


export class HotelRepository extends BaseRepository<Hotel> {
  constructor() {
    super(Hotel)
  }
  async findAll(): Promise<Hotel[]> {
    const hotels = await this.model.findAll({ where: { deletedAt: null } })
    if (!hotels) {
      return []
    }
    return hotels

  }
  async softDelete(id: number): Promise<Hotel | null> {
    const hotel = await this.model.findByPk(id)
    if (!hotel) {
      throw new NotFoundError("no hotel found")
    }
    hotel.deletedAt = new Date();
    await hotel.save();
    logger.info("hotel has been deleted")
    return hotel
  }
}