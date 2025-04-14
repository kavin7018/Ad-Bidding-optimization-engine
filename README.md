# Ad-Bidding-optimization-engine

Problem statement: There are multiple Adplatforms available in the market. As a client, I need a solution to identify which platform can maximize my ROI dynamically.

Design Solution:
1. I have a mock generator, which will generate events (view, click and conversion) per platform.
2. The mock server will publish events to Pulsar topics. This approach is highly reliable, as Pulsar is well-suited for handling high-throughput ingestion, making it more robust than directly pushing events to downstream systems.
3. Application will subscribe to those pulsar topics and receive those events. This platforms performance will be updated in real-time in the redis.
4. This approach of updating the state in redis instead of relation db which help to avoid latency issues and help in tackling heavy load situations.
5. This redis data will be periodically fetched, processed and aggregated every 10mins and stored in the database.
6. This data will be fetched by our AI module to do the performance analysis, the analytics module will provide the optimal time to bid, optimal amount and right platform to maximize the ROI.
7. This data will be fed as a input to our application periodically to improve the ROI.

Note: I have also created a grafana dashboard to visualize the real-time performance of the AdPlatforms.


Pending Items:
1. Aggregated metrics needs to be pushed to the relational database
2. AI module needs to be developed - Need to fetch the data and perform computations on top of it and provide the optimal bid timings, bid amount and ideal platform to bid.
3. Grafana json needs to be manually imported and data source needs to be configured manually - Need to automate [Mount the files to the grafana container]

Note: To view the grafana dashboard, import the dashboard.json and configure the prometheus data source[url:http://prometheus:9090]


![image](https://github.com/user-attachments/assets/5a3ed045-a8bb-47ac-b76f-ca4f6f7b99f9)
