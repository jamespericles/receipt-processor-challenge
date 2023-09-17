import { Router } from 'express'
import { processReceipt, getPoints } from '../controller'

const router: Router = Router()

router.post('/process', processReceipt)
router.get('/:id/points', getPoints)

export default router
