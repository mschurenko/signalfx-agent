from functools import partial as p
from textwrap import dedent

import pytest

from tests.helpers.agent import Agent
from tests.helpers.assertions import any_metric_found
from tests.helpers.metadata import Metadata
from tests.helpers.util import container_ip, get_monitor_metrics_from_selfdescribe, wait_for
from tests.helpers.verify import run_agent_verify

pytestmark = [pytest.mark.collectd, pytest.mark.openstack, pytest.mark.monitor_without_endpoints]


@pytest.mark.flaky(reruns=1)
def test_openstack(devstack):
    host = container_ip(devstack)
    config = dedent(
        f"""
            monitors:
              - type: collectd/openstack
                authURL: http://{host}/identity/v3
                username: admin
                password: testing123
        """
    )
    with Agent.run(config) as agent:
        expected_metrics = get_monitor_metrics_from_selfdescribe("collectd/openstack")
        assert wait_for(
            p(any_metric_found, agent.fake_services, expected_metrics), 60
        ), "Timed out waiting for openstack metrics"


METADATA = Metadata.from_package("collectd/openstack")
INCLUDED_METRICS = METADATA.included_metrics - {
    "gauge.openstack.nova.server.vda_write",
    "gauge.openstack.nova.server.memory-actual",
    "gauge.openstack.nova.server.vda_read_req",
    "gauge.openstack.nova.server.memory-rss",
    "gauge.openstack.nova.server.vda_write_req",
    "gauge.openstack.nova.server.memory",
    "gauge.openstack.nova.server.vda_read",
    # Openstack monitor does not emit any counters
    "counter.openstack.nova.server.rx",
    "counter.openstack.nova.server.rx_packets",
    "counter.openstack.nova.server.tx",
    "counter.openstack.nova.server.tx_packets",
}


@pytest.mark.flaky(reruns=1)
def test_openstacki_included(devstack):
    host = container_ip(devstack)
    run_agent_verify(
        f"""
            monitors:
            - type: collectd/openstack
              authURL: http://{host}/identity/v3
              username: admin
              password: testing123
        """,
        INCLUDED_METRICS,
    )
