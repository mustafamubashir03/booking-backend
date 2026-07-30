import Room from "../db/models/room";
import { BaseRepository } from "./base.repository";

export class RoomRepository extends BaseRepository<Room>{
    constructor(){
        super(Room)
    }
}