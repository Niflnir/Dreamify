from concurrent import futures
from openai import OpenAI
import os
import logging
import sys
import grpc
import image_generator_pb2
import image_generator_pb2_grpc
from grpc_reflection.v1alpha import reflection


class ImageGeneratorService(image_generator_pb2_grpc.ImageGenerator):
    def GetImageUrl(self, request, context):
        _logger.info(f"Request received from server: {request}")
        image_url = generateImageAndReturnImageUrl(request.prompt)
        _logger.info(f"Generated image url: {image_url}")
        return image_generator_pb2.ImageResponse(image_url=image_url)


def generateImageAndReturnImageUrl(prompt):
    response = client.images.generate(
      model="dall-e-3",
      prompt="Create a stunning wallpaper based on :[" + prompt + "]",
      size="1024x1024",
      quality="standard",
      n=1,
    )

    return response.data[0].url


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


def _init_logger():
    # Create a looger named 'app'
    logger = logging.getLogger('app')

    # Set the threshold logging level of the logger to INFO
    logger.setLevel(logging.INFO)

    # Create a stream-based handler that writes the log entries into the standard output stream
    handler = logging.StreamHandler(sys.stdout)

    # Create a formatter for the logs
    formatter = logging.Formatter('%(created)f:%(levelname)s:%(name)s:%(module)s:%(message)s')

    # Set the created formatter as the formatter of the handler    
    handler.setFormatter(formatter)

    # Add the created handler to this logger
    logger.addHandler(handler)


if __name__ == '__main__':
    _init_logger()
    _logger = logging.getLogger('app')

    # Initialize OpenAi client
    client = OpenAI(api_key=os.environ["OPENAI_API_KEY"])

    # Serve grpc server
    serve()
