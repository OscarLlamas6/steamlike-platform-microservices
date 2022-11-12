import { Router } from 'express'
import { addWishlist, getUserWishlist, updateWishlist, deleteWishlist, removeGameFromWishlist } from '../controllers/wishlist.controller'

const router = Router();

router.post("/create", addWishlist);
router.get("/:username", getUserWishlist);
router.put("/update", updateWishlist);
router.delete("/delete/:idMyGame", deleteWishlist);
router.delete("/remove/:idMyGame", removeGameFromWishlist);

export default router;