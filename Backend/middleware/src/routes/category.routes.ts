import { Router } from 'express'
import { addCategory,
         getCategory,
         getCategoriesList,
         updateCategory, 
         deleteCategory,
         getCategoriesListNoAuth
        } from '../controllers/category.controller'

const router = Router();

router.post("/create", addCategory);
router.get("/:idCategory", getCategory);
router.get("/list", getCategoriesList);
router.get("/list/noauth", getCategoriesListNoAuth);
router.put("/update", updateCategory);
router.delete("/delete/:idCategory", deleteCategory);

export default router;