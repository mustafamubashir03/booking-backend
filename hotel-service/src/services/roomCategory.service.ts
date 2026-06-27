import { createRoomCategoryDTO, updateRoomCategoryDTO } from "../dto/roomCategory.dto";
import { RoomCategoryRepository } from "../repositories/roomCategory.repository";


const roomCategoryRepository = new RoomCategoryRepository()


export const getRoomCategoryService = async () => {
    const roomCategories = await roomCategoryRepository.findAll()
    return roomCategories
}
export const createRoomCategoryService = async (roomCategoryData: createRoomCategoryDTO) => {
    const roomCategory = await roomCategoryRepository.create(roomCategoryData)
    return roomCategory
}
export const updateRoomCategoryService = async ({ id, roomCategoryData }: { id: number, roomCategoryData: updateRoomCategoryDTO }) => {
    const roomCategory = await roomCategoryRepository.update(id, roomCategoryData)
    return roomCategory
}
export const deleteRoomCategoryService = async (id: number) => {
    const roomCategory = await roomCategoryRepository.softDelete(id)
    return roomCategory
}