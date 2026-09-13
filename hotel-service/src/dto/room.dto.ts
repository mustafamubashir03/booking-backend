export interface RoomGenerationResponse {
    success: boolean;
    totalRoomsCreated: number,
    totalDatesCovered: number,
    errors: string[];
    jobId: string;
}
