import express from "express"

import { roomGenerationController } from "../../controllers/roomGeneration.controller"
import { validateRequestBody } from "../../validators";
import { RoomGenerationJobSchema } from "../../validators/roomGeneration.validator";


const roomRouter = express.Router();

roomRouter.post('/generate', validateRequestBody(RoomGenerationJobSchema), roomGenerationController)

export default roomRouter