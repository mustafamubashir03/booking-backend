import RoomCategory from "../db/models/roomCategory";
import { BaseRepository } from "./base.repository";

export class RoomCategoryRepository extends BaseRepository<RoomCategory> {
    constructor() {
        super(RoomCategory)
    }
    async findAll() {
        const roomCategories = await this.model.findAll({
            where: {
                deletedAt: null
            }
        })
        if (!roomCategories) {
            throw new Error("no room categories found")
        }
        return roomCategories
    }
    async softDelete(id: number): Promise<RoomCategory | null> {
        const roomCategory = await this.model.findByPk(id)
        if (!roomCategory) {
            throw new Error("no room category found")
        }
        roomCategory.deletedAt = new Date();
        await roomCategory.save();
        return roomCategory
    }

}