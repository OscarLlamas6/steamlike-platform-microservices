import { Router } from 'express'
import { addDiscount, getGamesDiscounts, getDLCDiscounts, getDiscountByID, updateDiscount, deleteDiscount } from '../controllers/discounts.controller'

const router = Router();

router.post("/create", addDiscount);
router.get("/single/:idDiscount", getDiscountByID);
router.get("/games", getGamesDiscounts);
router.get("/dlc", getDLCDiscounts);
router.put("/update", updateDiscount);
router.delete("/delete/:idDiscount", deleteDiscount);

export default router;