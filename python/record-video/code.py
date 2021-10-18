import cv2

webcam = 0
source = cv2.VideoCapture(webcam)
has_frame, frame = source.read()
if not has_frame:
    print('no frame')
    raise ValueError('No Webcam')

width = int(source.get(cv2.CAP_PROP_FRAME_WIDTH))
height = int(source.get(cv2.CAP_PROP_FRAME_HEIGHT))

out = cv2.VideoWriter('test.mp4', cv2.VideoWriter_fourcc(*'MJPG'), 10, (width, height))
win = cv2.namedWindow('WebCam', cv2.WINDOW_NORMAL)
while cv2.waitKey(1) != 27:
    has_frame, frame = source.read()
    if not has_frame:
        break
    cv2.imshow(win, frame)
    out.write(frame)
source.release()
out.release()
cv2.destroyWindow(win)
