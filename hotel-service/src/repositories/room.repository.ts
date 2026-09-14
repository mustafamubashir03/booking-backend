import { CreationAttributes, Op } from "sequelize";
import Room from "../db/models/room";
import { BaseRepository } from "./base.repository";

export class RoomRepository extends BaseRepository<Room> {
    constructor() {
        super(Room)
    }
    async findByRoomCategoryIdAndDateRange(roomCategoryId: number, start: Date, end: Date) {
        return await this.model.findAll({
            where: {
                roomCategoryId: roomCategoryId,
                dateOfAvailability: {
                    [Op.between]: [start, end]
                },
                deletedAt: null
            }
        })
    }

    async findLatestByRoomCategoryId(roomCategoryId: number) {
        return await this.model.findOne({
            where: {
                roomCategoryId,
                deletedAt: null
            },
            order: [['dateOfAvailability', 'DESC']],
        });
    }
    async bulkCreate(rooms: CreationAttributes<Room>[]) {
        return await this.model.bulkCreate(rooms)
    }
}