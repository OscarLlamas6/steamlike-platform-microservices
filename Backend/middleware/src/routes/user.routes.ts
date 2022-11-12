import { Router } from 'express'
import { createUser, updateUser, deleteUser, getProfile, getAllUsers, loginUser, verifyUser, resendVerify, restoreUser } from '../controllers/user.controller'

const router = Router();

router.get("/profile/:username", getProfile);
router.get("/list/all", getAllUsers);
router.post("/create", createUser);
router.post("/login", loginUser);
router.put("/update", updateUser);
router.delete("/delete/:username", deleteUser);
router.get("/restore/:username", restoreUser);
router.get("/verify", verifyUser);
router.get("/resend/verify", resendVerify);

export default router;