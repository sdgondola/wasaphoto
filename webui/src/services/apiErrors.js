
export const InternalServerError = new Error("Internal server error");
export const BadIdsException = new Error("Bad auth token or bad userID");
export const BlockedException = new Error("Forbidden: user blocked you!");
export const UserNotFoundException = new Error("User not found");