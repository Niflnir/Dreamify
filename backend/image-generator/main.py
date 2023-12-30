from concurrent import futures
from openai import OpenAI
import grpc
import image_generator_pb2
import image_generator_pb2_grpc
from grpc_reflection.v1alpha import reflection

# client = OpenAI(api_key="")
# response = client.images.generate(
#   model="dall-e-3",
#   prompt="a white siamese cat",
#   size="1024x1024",
#   quality="standard",
#   n=1,
# )
#
# image_url = response.data[0].url

class ImageGeneratorService(image_generator_pb2_grpc.ImageGenerator):
    def GetImageUrl(self, request, context):
        image_url = "test url"
        return image_generator_pb2.ImageResponse(image_url=image_url)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    image_generator_pb2_grpc.add_ImageGeneratorServicer_to_server(ImageGeneratorService(), server)
    SERVICE_NAMES = (
            image_generator_pb2.DESCRIPTOR.services_by_name['ImageGenerator'].full_name,
            reflection.SERVICE_NAME
    )
    reflection.enable_server_reflection(SERVICE_NAMES, server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print("Image generator service started!")
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
