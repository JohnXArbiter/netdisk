export interface Share {
    id: number
    file: {
        fileId: number
        name: string
    }
    description: string;
    created_at: string;
    updated_at: string;
}